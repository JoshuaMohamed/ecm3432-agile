import "./StarRating.css";

type StarRatingProps = {
  rating: number;
  reviews: number;
};

function StarRating({ rating, reviews }: StarRatingProps) {
  return (
    <div className="rating">
      {[0, 1, 2, 3, 4].map((index) => (
        <span
          key={index}
          className={`star ${index < rating ? "filled" : ""}`}
        />
      ))}
      <span className="reviews">{reviews} reviews</span>
    </div>
  );
}

export default StarRating;
