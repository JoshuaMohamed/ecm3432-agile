import "./PlaceCard.css";

type PlaceCardProps = {
  name: string;
  postcode: string;
  rating?: number;
  reviews?: number;
  summary?: string;
};

function PlaceCard({
  name,
  postcode,
  rating,
  reviews,
  summary,
}: PlaceCardProps) {
  const safeRating = rating ?? 0;
  const safeReviews = reviews ?? 0;
  const safeSummary = summary ?? "No description yet.";

  return (
    <article className="card">
      <div className="thumb" />
      <div className="details">
        <div className="title-row">
          <div>
            <h2>{name}</h2>
            <p className="postcode">{postcode}</p>
          </div>
          <div className="rating">
            {[0, 1, 2, 3, 4].map((index) => (
              <span
                key={index}
                className={`star ${index < safeRating ? "filled" : ""}`}
              />
            ))}
            <span className="reviews">{safeReviews} reviews</span>
          </div>
        </div>
        <p className="summary">{safeSummary}</p>
      </div>
    </article>
  );
}

export default PlaceCard;
