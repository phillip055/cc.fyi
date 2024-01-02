json.extract! short_url, :id, :long_url, :hash_key, :created_at, :updated_at
json.url short_url_url(short_url, format: :json)
