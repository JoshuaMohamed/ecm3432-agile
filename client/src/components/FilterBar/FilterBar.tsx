import { useState, useRef, useEffect } from "react";
import "./FilterBar.css";

function FilterBar() {
  const options = ["Area", "District", "Sector"] as const;

  const [selected, setSelected] = useState<(typeof options)[number]>("Area");
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
    setActiveIndex(options.indexOf(selected));
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
      setActiveIndex((index) => (index + 1) % options.length);
    }

    if (event.key === "ArrowUp") {
      event.preventDefault();
      setActiveIndex((index) => (index - 1 + options.length) % options.length);
    }

    if (event.key === "Enter") {
      event.preventDefault();
      setSelected(options[activeIndex]);
      setIsOpen(false);
    }
  };

  return (
    <section className="filters">
      <label className="search">
        <span className="sr-only">Postcode</span>
        <input type="text" placeholder="Enter Postcode..." />
        <span className="search-icon" aria-hidden="true" />
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
          {selected}
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
          {options.map((option, index) => (
            <li
              key={option}
              id={`option-${index}`}
              role="option"
              aria-selected={selected === option}
            >
              <button
                type="button"
                className={`select-option ${activeIndex === index ? "active" : ""}`}
                onClick={() => {
                  setSelected(option);
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
    </section>
  );
}

export default FilterBar;
