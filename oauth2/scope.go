package oauth

type Scope string

// List of Scopes
var Scopes = struct {
	ENTITY_CLIENTS_READ                  Scope
	ENTITY_CLIENTS_ALL                   Scope
	ENTITY_SUPPLIERS_READ                Scope
	ENTITY_SUPPLIERS_ALL                 Scope
	PRODUCTS_READ                        Scope
	PRODUCTS_ALL                         Scope
	ISSUED_DOCUMENTS_INVOICES_READ       Scope
	ISSUED_DOCUMENTS_CREDIT_NOTES_READ   Scope
	ISSUED_DOCUMENTS_RECEIPTS_READ       Scope
	ISSUED_DOCUMENTS_ORDERS_READ         Scope
	ISSUED_DOCUMENTS_QUOTES_READ         Scope
	ISSUED_DOCUMENTS_PROFORMAS_READ      Scope
	ISSUED_DOCUMENTS_DELIVERY_NOTES_READ Scope
	ISSUED_DOCUMENTS_INVOICES_ALL        Scope
	ISSUED_DOCUMENTS_CREDIT_NOTES_ALL    Scope
	ISSUED_DOCUMENTS_RECEIPTS_ALL        Scope
	ISSUED_DOCUMENTS_ORDERS_ALL          Scope
	ISSUED_DOCUMENTS_QUOTES_ALL          Scope
	ISSUED_DOCUMENTS_PROFORMAS_ALL       Scope
	ISSUED_DOCUMENTS_DELIVERY_NOTES_ALL  Scope
	RECEIVED_DOCUMENTS_READ              Scope
	RECEIVED_DOCUMENTS_ALL               Scope
	STOCK_READ                           Scope
	STOCK_ALL                            Scope
	RECEIPTS_READ                        Scope
	RECEIPTS_ALL                         Scope
	TAXES_READ                           Scope
	TAXES_ALL                            Scope
	ARCHIVE_READ                         Scope
	ARCHIVE_ALL                          Scope
	CASHBOOK_READ                        Scope
	CASHBOOK_ALL                         Scope
	SETTINGS_READ                        Scope
	SETTINGS_ALL                         Scope
	SITUATION_READ                       Scope
}{
	ENTITY_CLIENTS_READ:                  "entity.clients:r",
	ENTITY_CLIENTS_ALL:                   "entity.clients:a",
	ENTITY_SUPPLIERS_READ:                "entity.suppliers:r",
	ENTITY_SUPPLIERS_ALL:                 "entity.suppliers:a",
	PRODUCTS_READ:                        "products:r",
	PRODUCTS_ALL:                         "products:a",
	ISSUED_DOCUMENTS_INVOICES_READ:       "issued_documents.invoices:r",
	ISSUED_DOCUMENTS_CREDIT_NOTES_READ:   "issued_documents.credit_notes:r",
	ISSUED_DOCUMENTS_RECEIPTS_READ:       "issued_documents.receipts:r",
	ISSUED_DOCUMENTS_ORDERS_READ:         "issued_documents.orders:r",
	ISSUED_DOCUMENTS_QUOTES_READ:         "issued_documents.quotes:r",
	ISSUED_DOCUMENTS_PROFORMAS_READ:      "issued_documents.proformas:r",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_READ: "issued_documents.delivery_notes:r",
	ISSUED_DOCUMENTS_INVOICES_ALL:        "issued_documents.invoices:a",
	ISSUED_DOCUMENTS_CREDIT_NOTES_ALL:    "issued_documents.credit_notes:a",
	ISSUED_DOCUMENTS_RECEIPTS_ALL:        "issued_documents.receipts:a",
	ISSUED_DOCUMENTS_ORDERS_ALL:          "issued_documents.orders:a",
	ISSUED_DOCUMENTS_QUOTES_ALL:          "issued_documents.quotes:a",
	ISSUED_DOCUMENTS_PROFORMAS_ALL:       "issued_documents.proformas:a",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_ALL:  "issued_documents.delivery_notes:a",
	RECEIVED_DOCUMENTS_READ:              "received_documents:r",
	RECEIVED_DOCUMENTS_ALL:               "received_documents:a",
	STOCK_READ:                           "stock:r",
	STOCK_ALL:                            "stock:a",
	RECEIPTS_READ:                        "receipts:r",
	RECEIPTS_ALL:                         "receipts:a",
	TAXES_READ:                           "taxes:r",
	TAXES_ALL:                            "taxes:a",
	ARCHIVE_READ:                         "archive:r",
	ARCHIVE_ALL:                          "archive:a",
	CASHBOOK_READ:                        "cashbook:r",
	CASHBOOK_ALL:                         "cashbook:a",
	SETTINGS_READ:                        "settings:r",
	SETTINGS_ALL:                         "settings:a",
	SITUATION_READ:                       "situation:r",
}
