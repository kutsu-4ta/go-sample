package repositories

type IItemRepository interface {
        FindAll(*[]models.Item, error)
}

type IItemMemoryRepository struct {
    item []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
    return & NewItemMemoryRepository{items: items}
}

func (r *NewItemMemoryRepository) FindAll() (*[]models.Item, error) {
    return &r.items, nil
}