import { useNavigate } from "react-router-dom";
import "./TopBar.css";

function TopBar() {
  const navigate = useNavigate();

  return (
    <header className="topbar">
      <h1 className="title" onClick={() => navigate("/")} style={{ cursor: "pointer" }}>
        Community Tourist Assistant
      </h1>
      <button
        className="profile-button"
        aria-label="User profile"
        onClick={() => navigate("/account")}
      >
        <span className="profile-icon" />
      </button>
    </header>
  );
}

export default TopBar;
