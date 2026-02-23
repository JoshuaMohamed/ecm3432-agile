import { useState } from "react";

import TopBar from "../../components/TopBar/TopBar";
import FilterBar from "../../components/FilterBar/FilterBar";
import PlaceCard from "../../components/PlaceCard/PlaceCard";
import { filterParam, type FilterOption } from "../../data/filters";
import "./PlacesPage.css";

function PlacesPage() {
  const [postcode, setPostcode] = useState("");
  const [selectedFilter, setSelectedFilter] = useState<FilterOption>("Area");
  const [places, setPlaces] = useState<{ name: string; postcode: string }[]>(
    [],
  );
  const [error, setError] = useState<string | null>(null);

  const handleSearch = async () => {
    setError(null);

    const url = new URL("http://localhost:8080/getPlaces");
    url.searchParams.set("postcode", postcode);
    url.searchParams.set("filter", filterParam(selectedFilter));

    try {
      const response = await fetch(url.toString());
      const data = (await response.json()) as {
        Data: { name: string; postcode: string }[];
        Message?: string;
      };

      if (!response.ok) {
        throw new Error(data.Message ?? "Request failed");
      }

      setPlaces(data.Data ?? []);
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Something went wrong";
      setError(message);
      setPlaces([]);
    }
  };

  return (
    <div className="page">
      <TopBar />

      <main className="content">
        <FilterBar
          postcode={postcode}
          onPostcodeChange={setPostcode}
          selectedFilter={selectedFilter}
          onFilterChange={setSelectedFilter}
          onSubmit={handleSearch}
        />

        {error && <p className="error">{error}</p>}

        <section className="list">
          {places.map((place) => (
            <PlaceCard
              key={`${place.name}-${place.postcode}`}
              name={place.name}
              postcode={place.postcode}
              rating={0}
              reviews={0}
              summary="No description yet."
            />
          ))}
        </section>
      </main>
    </div>
  );
}

export default PlacesPage;
