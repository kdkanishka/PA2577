class ShoppingItem {
  final String id;
  final String name;
  final bool isCompleted;
  final String shoppingListId;

  ShoppingItem({
    required this.id,
    required this.name,
    required this.isCompleted,
    required this.shoppingListId,
  });

  factory ShoppingItem.fromJson(Map<String, dynamic> json) {
    return ShoppingItem(
      id: json['id'],
      name: json['name'],
      isCompleted: json['completed'] ?? false,
      shoppingListId: json['shopping_list_id'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'completed': isCompleted,
      'shopping_list_id': shoppingListId,
    };
  }
}