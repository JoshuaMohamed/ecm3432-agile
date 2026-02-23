import { useState, useRef, useEffect } from "react";
import { filterOptions, type FilterOption } from "../../data/filters";
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
  const [isOpen, setIsOpen] = useState(false);
  const [activeIndex, setActiveIndex] = useState(0);

  const selectRef = useRef<HTMLDivElement>(null);
  const menuRef = useRef<HTMLUListElement>(null);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (!selectRef.current?.contains(event.target as Node)) {
        setIsOpen(false);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  });

  useEffect(() => {
    if (isOpen) {
      menuRef.current?.focus();
    }
  }, [isOpen]);

  const openMenu = () => {
    setActiveIndex(filterOptions.indexOf(selectedFilter));
    setIsOpen(true);
  };

  const handleTriggerKeyDown = (event: React.KeyboardEvent) => {
    if (event.key === "ArrowDown" || event.key === "ArrowUp") {
      event.preventDefault();
      openMenu();
    }

    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      setIsOpen((open) => !open);
    }
  };

  const handleMenuKeyDown = (event: React.KeyboardEvent) => {
    if (event.key === "Escape") {
      event.preventDefault();
      setIsOpen(false);
    }

    if (event.key === "ArrowDown") {
      event.preventDefault();
      setActiveIndex((index) => (index + 1) % filterOptions.length);
    }

    if (event.key === "ArrowUp") {
      event.preventDefault();
      setActiveIndex(
        (index) => (index - 1 + filterOptions.length) % filterOptions.length,
      );
    }

    if (event.key === "Enter") {
      event.preventDefault();
      onFilterChange(filterOptions[activeIndex]);
      setIsOpen(false);
    }
  };

  const handleSubmit = (event: React.SubmitEvent) => {
    event.preventDefault();
    onSubmit();
  };

  return (
    <form className="filters" onSubmit={handleSubmit}>
      <label className="search">
        <span className="sr-only">Postcode</span>
        <input
          type="text"
          placeholder="Enter Postcode..."
          // value + onChange makes the input controlled
          value={postcode}
          onChange={(event) => onPostcodeChange(event.target.value)}
        />
        <button className="search-button" type="submit" aria-label="Search">
          <span className="search-icon" aria-hidden="true" />
        </button>
      </label>

      <div className="select" ref={selectRef}>
        <button
          className="select-trigger"
          type="button"
          onClick={() => {
            isOpen ? setIsOpen(false) : openMenu();
          }}
          onKeyDown={handleTriggerKeyDown}
          aria-expanded={isOpen}
          aria-haspopup="listbox"
        >
          {selectedFilter}
          <span className="chevron" aria-hidden="true" />
        </button>

        <ul
          className={`select-menu ${isOpen ? "open" : ""}`}
          role="listbox"
          aria-activedescendant={`option-${activeIndex}`}
          tabIndex={-1}
          ref={menuRef}
          onKeyDown={handleMenuKeyDown}
        >
          {filterOptions.map((option, index) => (
            <li
              key={option}
              id={`option-${index}`}
              role="option"
              aria-selected={selectedFilter === option}
            >
              <button
                type="button"
                className={`select-option ${activeIndex === index ? "active" : ""}`}
                onClick={() => {
                  onFilterChange(option);
                  setActiveIndex(index);
                  setIsOpen(false);
                }}
              >
                {option}
              </button>
            </li>
          ))}
        </ul>
      </div>
    </form>
  );
}

export default FilterBar;
