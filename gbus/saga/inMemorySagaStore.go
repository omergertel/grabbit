package saga

import (
	"database/sql"
	"errors"
)

type InMemorySagaStore struct {
	instances map[*SagaDef][]*SagaInstance
}

func (store *InMemorySagaStore) GetSagaByID(tx *sql.Tx, sagaID string) (*SagaInstance, error) {
	for _, instances := range store.instances {
		for _, instance := range instances {
			if instance.ID == sagaID {
				return instance, nil
			}
		}
	}
	return nil, errors.New("no saga found for provided id")
}
func (store *InMemorySagaStore) SaveNewSaga(tx *sql.Tx, def *SagaDef, newInstance *SagaInstance) error {
	instances := store.instances[def]
	if instances == nil {
		instances = make([]*SagaInstance, 0)

	}
	instances = append(instances, newInstance)
	store.instances[def] = instances
	return nil

}
func (store *InMemorySagaStore) UpdateSaga(tx *sql.Tx, instance *SagaInstance) error {

	return nil
}

func NewInMemoryStore() SagaStore {
	return &InMemorySagaStore{
		instances: make(map[*SagaDef][]*SagaInstance)}
}
