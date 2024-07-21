package structs


type CacheItem struct {
    Key        string
    Value      interface{}
    Expiration int64
}

type SetRequest struct {
    Key        string      `json:"key"`
    Value      interface{} `json:"value"`
    Expiration int         `json:"expiration"` // in seconds
}

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
