import TopBar from "../../components/TopBar/TopBar";
import FilterBar from "../../components/FilterBar/FilterBar";
import PlaceCard from "../../components/PlaceCard/PlaceCard";
import { mockPlaces } from "../../data/mockPlaces";
import "./PlacesPage.css";

function PlacesPage() {
  return (
    <div className="page">
      <TopBar />

      <main className="content">
        <FilterBar />

        <section className="list">
          {mockPlaces.map((place) => (
            <PlaceCard key={place.name} {...place} />
          ))}
        </section>
      </main>
    </div>
  );
}

export default PlacesPage;
