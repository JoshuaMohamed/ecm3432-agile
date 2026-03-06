import { Routes, Route } from "react-router-dom";
import PlacesPage from "../pages/Places/PlacesPage";
import AccountPage from "../pages/Account/AccountPage";

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<PlacesPage />} />
      <Route path="/account" element={<AccountPage />} />
    </Routes>
  );
}
