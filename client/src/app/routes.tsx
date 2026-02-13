import { Routes, Route } from "react-router-dom";
import PlacesPage from "../pages/Places/PlacesPage";

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<PlacesPage />} />
    </Routes>
  );
}
