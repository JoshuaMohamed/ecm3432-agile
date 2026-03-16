import "./PostcodeSearch.css";

type PostcodeSearchProps = {
  postcode: string;
  onPostcodeChange: (value: string) => void;
};

function PostcodeSearch({ postcode, onPostcodeChange }: PostcodeSearchProps) {
  return (
    <label className="search">
      <span className="sr-only">Postcode</span>
      <input
        type="text"
        placeholder="Enter Postcode..."
        value={postcode}
        onChange={(event) => onPostcodeChange(event.target.value)}
      />
      <button className="search-button" type="submit" aria-label="Search">
        <span className="search-icon" aria-hidden="true" />
      </button>
    </label>
  );
}

export default PostcodeSearch;
