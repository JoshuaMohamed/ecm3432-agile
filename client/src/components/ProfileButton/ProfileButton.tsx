import "./ProfileButton.css";

type ProfileButtonProps = {
  onClick: () => void;
};

function ProfileButton({ onClick }: ProfileButtonProps) {
  return (
    <button
      className="profile-button"
      aria-label="User profile"
      onClick={onClick}
    >
      <span className="profile-icon" />
    </button>
  );
}

export default ProfileButton;
