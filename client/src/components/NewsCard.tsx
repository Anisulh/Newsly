type NewsItem = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Title: string;
  Description: string;
  Content: string;
  URL: string;
  ImageURL: string;
  PublishedAt: string;
  Source: string;
  Category: string;
  Keywords: string[];
  Likes: number | null;
  Dislikes: number | null;
};

export default function NewsCard({ news }: { news: NewsItem }) {
  return (
    <div className="card">
      <img src={news.ImageURL} className="card-img-top" alt={news.Title} />
      <div className="card-body">
        <div className="category-div">
          <small className="category">{news.Category}</small>
        </div>

        <h5 className="card-title">{news.Title}</h5>
        <p className="card-text">{news.Description}</p>
        <p className="card-text">
          
        </p>

        <a href={news.URL} className="btn btn-primary">
          Read more
        </a>
      </div>
      <div className="card-footer">
        <small className="text-muted">Source: {news.Source}</small>
        <small className="text-muted">
            Published: {new Date(news.PublishedAt).toLocaleDateString()}
          </small>
      </div>
    </div>
  );
}
