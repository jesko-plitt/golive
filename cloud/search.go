package cloud

import (
	"strconv"

	"github.com/ao-concepts/logging"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func ProvideTypesenseClient(log logging.Logger) *typesense.Client {
	cfg := ProvideTypesenseConfig()

	return typesense.NewClient(
		typesense.WithServer(cfg.Server),
		typesense.WithAPIKey(cfg.ApiKey))
}

type SearchIndex struct {
	typesenseClient *typesense.Client
}

func ProvideSearchIndex(log logging.Logger) *SearchIndex {
	return &SearchIndex{
		typesenseClient: ProvideTypesenseClient(log),
	}
}

type Indexable interface {
	GetID() uint
	GetIndexConfig(c interface{}) *api.CollectionSchema
	ToDocument(c interface{}) (document map[string]interface{}, err error)
}

func (si *SearchIndex) Add(entity Indexable, container interface{}) error {
	document, err := entity.ToDocument(container)

	if err != nil {
		return err
	}

	_, err = si.typesenseClient.
		Collection(entity.GetIndexConfig(container).Name).
		Documents().
		Upsert(document)

	return err
}

func (si *SearchIndex) Remove(entity Indexable, container interface{}) error {
	document := si.typesenseClient.
		Collection(entity.GetIndexConfig(container).Name).
		Document(
			strconv.FormatUint(uint64(entity.GetID()), 10),
		)

	_, err := document.Retrieve()

	if err == nil {
		_, err := document.Delete()
		return err
	}

	return nil
}
