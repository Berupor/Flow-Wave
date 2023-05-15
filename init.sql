CREATE TABLE IF NOT EXISTS reviews
(
    ProductID Int32,
    PlaceID   Int32,
    AuthorID  Int32,
    Rating    Float64,
    Review    String,
    Timestamp DateTime
) ENGINE = MergeTree()
      ORDER BY Timestamp;
