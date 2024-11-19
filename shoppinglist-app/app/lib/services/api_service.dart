import 'package:http/http.dart' as http;
import 'dart:convert';
import '../models/shopping_list.dart';
import '../models/shopping_item.dart';

class ApiService {
  final String baseUrl = 'http://192.168.49.2:30080'; // Empty for relative URLs
  final username = 'test';
  final password = 'test';

  Map<String, String> _createAuthHeader() {
    final credentials = base64Encode(utf8.encode('$username:$password'));
    final headers = {
      'Authorization': 'Basic $credentials',
    };
    return headers;
  }

  // Shopping Lists API calls
  Future<List<ShoppingList>> getShoppingLists() async {
    Map<String, String> headers = _createAuthHeader();

    final response = await http.get(
      Uri.parse('$baseUrl/shoppinglists'),''
      headers: headers,
    );

    if (response.statusCode == 200) {
      final List<dynamic> jsonList = json.decode(response.body);
      return jsonList.map((json) => ShoppingList.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load shopping lists');
    }
  }

  Future<ShoppingList> createShoppingList(String name) async {
    Map<String, String> headers= _createAuthHeader();
    headers['Content-Type'] = 'application/json';

    final response = await http.post(
      Uri.parse('$baseUrl/shoppinglists'),
      headers: headers,
      body: json.encode({'name': name}),
    );
    if (response.statusCode == 201) {
      return ShoppingList.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to create shopping list');
    }
  }

  Future<List<ShoppingItem>> getShoppingItems(String listId) async {
    Map<String, String> headers= _createAuthHeader();
    final response = await http.get(
      Uri.parse('$baseUrl/shoppinglists/$listId/shoppingitems'),
      headers: headers,
    );
    if (response.statusCode == 200) {
      final List<dynamic> jsonList = json.decode(response.body);
      return jsonList.map((json) => ShoppingItem.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load shopping items');
    }
  }

  Future<ShoppingItem> createShoppingItem(String name, String listId) async {
    Map<String, String> headers= _createAuthHeader();
    headers['Content-Type'] = 'application/json';

    final response = await http.post(
      Uri.parse('$baseUrl/shoppingitems'),
      headers: headers,
      body: json.encode({
        'name': name,
        'shopping_list_id': listId,
      }),
    );
    if (response.statusCode == 201) {
      return ShoppingItem.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to create shopping item');
    }
  }

  Future<void> completeShoppingItem(String itemId) async {
    Map<String, String> headers= _createAuthHeader();
    headers['Content-Type'] = 'application/json';

    final response = await http.put(
      Uri.parse('$baseUrl/shoppingitems/$itemId/complete'),
      headers: headers,
    );
    if (response.statusCode != 200) {
      throw Exception('Failed to complete shopping item');
    }
  }
}
