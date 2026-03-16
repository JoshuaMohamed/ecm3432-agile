import { useNavigate } from "react-router-dom";
import ProfileButton from "../ProfileButton/ProfileButton";
import "./TopBar.css";

function TopBar() {
  const navigate = useNavigate();

  return (
    <header className="topbar">
      <h1
        className="title"
        onClick={() => navigate("/")}
        style={{ cursor: "pointer" }}
      >
        Community Tourist Assistant
      </h1>
      <ProfileButton onClick={() => navigate("/account")} />
    </header>
  );
}

export default TopBar;
