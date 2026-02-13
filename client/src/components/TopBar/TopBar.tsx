import "./TopBar.css"

function TopBar() {
  return (
    <header className="topbar">
      <h1 className="title">Community Tourist Assistant</h1>
      <button className="profile-button" aria-label="User profile">
        <span className="profile-circle" />
      </button>
    </header>
  );
}

export default TopBar;
