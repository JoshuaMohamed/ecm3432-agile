import { type FilterOption } from "../../data/filters";
import PostcodeSearch from "../PostcodeSearch/PostcodeSearch";
import FilterSelect from "../FilterSelect/FilterSelect";
import "./FilterBar.css";

// Dev note: controlled inputs keep form values in React state so parent can submit them
type FilterBarProps = {
  postcode: string;
  onPostcodeChange: (value: string) => void;
  selectedFilter: FilterOption;
  onFilterChange: (value: FilterOption) => void;
  onSubmit: () => void;
};

function FilterBar({
  postcode,
  onPostcodeChange,
  selectedFilter,
  onFilterChange,
  onSubmit,
}: FilterBarProps) {
  const handleSubmit = (event: React.SubmitEvent) => {
    event.preventDefault();
    onSubmit();
  };

  return (
    <form className="filters" onSubmit={handleSubmit}>
      <PostcodeSearch postcode={postcode} onPostcodeChange={onPostcodeChange} />
      <FilterSelect
        selectedFilter={selectedFilter}
        onFilterChange={onFilterChange}
      />
    </form>
  );
}

export default FilterBar;
