import "./PlaceCard.css";

type PlaceCardProps = {
  name: string;
  postcode: string;
  rating: number;
  reviews: number;
  summary: string;
};

function PlaceCard({
  name,
  postcode,
  rating,
  reviews,
  summary,
}: PlaceCardProps) {
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
                className={`star ${index < rating ? "filled" : ""}`}
              />
            ))}
            <span className="reviews">{reviews} reviews</span>
          </div>
        </div>
        <p className="summary">{summary}</p>
      </div>
    </article>
  );
}

export default PlaceCard;
