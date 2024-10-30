class ShoppingList {
  final String id;
  final String name;
  final String description;
  
  ShoppingList({
    required this.id,
    required this.name,
    required this.description,
  });

  factory ShoppingList.fromJson(Map<String, dynamic> json) {
    return ShoppingList(
      id: json['id'],
      name: json['name'],
      description: json['description'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description' : description,
    };
  }
}