import 'package:flutter/material.dart';
import '../services/api_service.dart';
import '../models/shopping_list.dart';
import '../models/shopping_item.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final ApiService _apiService = ApiService();
  List<ShoppingList> _shoppingLists = [];
  Map<String, List<ShoppingItem>> _shoppingItems = {};
  String? _selectedListId;
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _loadShoppingLists();
  }

  Future<void> _loadShoppingLists() async {
    setState(() => _isLoading = true);
    try {
      final lists = await _apiService.getShoppingLists();
      setState(() => _shoppingLists = lists);
    } catch (e) {
      _showError('Failed to load shopping lists: $e');
    } finally {
      setState(() => _isLoading = false);
    }
  }

  Future<void> _loadShoppingItems(String listId) async {
    try {
      final items = await _apiService.getShoppingItems(listId);
      setState(() => _shoppingItems[listId] = items);
    } catch (e) {
      _showError('Failed to load shopping items: $e');
    }
  }

  void _showError(String message) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text(message), backgroundColor: Colors.red),
    );
  }

  Future<void> _addNewShoppingList() async {
    final nameController = TextEditingController();
    return showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text('New Shopping List'),
        content: TextField(
          controller: nameController,
          decoration: InputDecoration(hintText: 'Enter list name'),
          autofocus: true,
        ),
        actions: [
          TextButton(
            child: Text('Cancel'),
            onPressed: () => Navigator.of(context).pop(),
          ),
          TextButton(
            child: Text('Create'),
            onPressed: () async {
              if (nameController.text.isNotEmpty) {
                try {
                  await _apiService.createShoppingList(nameController.text);
                  Navigator.of(context).pop();
                  _loadShoppingLists();
                } catch (e) {
                  _showError('Failed to create list: $e');
                }
              }
            },
          ),
        ],
      ),
    );
  }

  Future<void> _addNewShoppingItem(String listId) async {
    final nameController = TextEditingController();
    return showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text('New Item'),
        content: TextField(
          controller: nameController,
          decoration: InputDecoration(hintText: 'Enter item name'),
          autofocus: true,
        ),
        actions: [
          TextButton(
            child: Text('Cancel'),
            onPressed: () => Navigator.of(context).pop(),
          ),
          TextButton(
            child: Text('Add'),
            onPressed: () async {
              if (nameController.text.isNotEmpty) {
                try {
                  await _apiService.createShoppingItem(nameController.text, listId);
                  Navigator.of(context).pop();
                  _loadShoppingItems(listId);
                } catch (e) {
                  _showError('Failed to create item: $e');
                }
              }
            },
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Shopping Lists'),
        actions: [
          IconButton(
            icon: Icon(Icons.add),
            onPressed: _addNewShoppingList,
          ),
        ],
      ),
      body: _isLoading
          ? Center(child: CircularProgressIndicator())
          : Row(
              children: [
                // Shopping Lists Panel
                Container(
                  width: 300,
                  decoration: BoxDecoration(
                    border: Border(
                      right: BorderSide(color: Colors.grey.shade300),
                    ),
                  ),
                  child: ListView.builder(
                    itemCount: _shoppingLists.length,
                    itemBuilder: (context, index) {
                      final list = _shoppingLists[index];
                      return ListTile(
                        title: Text(list.name),
                        selected: _selectedListId == list.id,
                        onTap: () {
                          setState(() => _selectedListId = list.id);
                          _loadShoppingItems(list.id);
                        },
                      );
                    },
                  ),
                ),
                // Shopping Items Panel
                Expanded(
                  child: _selectedListId == null
                      ? Center(child: Text('Select a shopping list'))
                      : Column(
                          children: [
                            AppBar(
                              title: Text('Items'),
                              automaticallyImplyLeading: false,
                              actions: [
                                IconButton(
                                  icon: Icon(Icons.add),
                                  onPressed: () => _addNewShoppingItem(_selectedListId!),
                                ),
                              ],
                            ),
                            Expanded(
                              child: _shoppingItems[_selectedListId] == null
                                  ? Center(child: CircularProgressIndicator())
                                  : ListView.builder(
                                      itemCount: _shoppingItems[_selectedListId]!.length,
                                      itemBuilder: (context, index) {
                                        final item = _shoppingItems[_selectedListId]![index];
                                        return ListTile(
                                          title: Text(
                                            item.name,
                                            style: TextStyle(
                                              decoration: item.isCompleted
                                                  ? TextDecoration.lineThrough
                                                  : null,
                                            ),
                                          ),
                                          trailing: Checkbox(
                                            value: item.isCompleted,
                                            onChanged: (bool? value) async {
                                              try {
                                                await _apiService.completeShoppingItem(item.id);
                                                _loadShoppingItems(_selectedListId!);
                                              } catch (e) {
                                                _showError('Failed to update item: $e');
                                              }
                                            },
                                          ),
                                        );
                                      },
                                    ),
                            ),
                          ],
                        ),
                ),
              ],
            ),
    );
  }
}