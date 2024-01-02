class CreateShortUrls < ActiveRecord::Migration[7.1]
  def change
    create_table :short_urls do |t|
      t.string :long_url
      t.string :hash_key, null: false, index: { unique: true }

      t.timestamps
    end
  end
end
