import StarRating from "../StarRating/StarRating";
import "./PlaceCard.css";

type PlaceCardProps = {
  name: string;
  postcode: string;
  coverPath: string;
  rating?: number;
  reviews?: number;
  summary?: string;
};

function PlaceCard({
  name,
  postcode,
  coverPath,
  rating,
  reviews,
  summary,
}: PlaceCardProps) {
  const safeRating = rating ?? 0;
  const safeReviews = reviews ?? 0;
  const safeSummary = summary ?? "No description yet.";
  const imageServer = "http://localhost:8080/assets/";

  return (
    <article className="card">
      <img src={`${imageServer}${coverPath}`} className="cover" />
      <div className="details">
        <div className="title-row">
          <div>
            <h2>{name}</h2>
            <p className="postcode">{postcode}</p>
          </div>
          <StarRating rating={safeRating} reviews={safeReviews} />
        </div>
        <p className="summary">{safeSummary}</p>
      </div>
    </article>
  );
}

export default PlaceCard;
