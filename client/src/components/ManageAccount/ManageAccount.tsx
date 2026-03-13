import "../AccountForm.css";

interface ManageAccountProps {
  email?: string | null;
  onLogOut: () => void;
}

function ManageAccount({ email, onLogOut }: ManageAccountProps) {
  return (
    <div className="account-manage">
      <h2 className="account-heading manage-heading">Manage Account</h2>
      {email && (
        <p className="account-email">
          <strong>Email:</strong> {email}
        </p>
      )}
      <button className="logout-button" onClick={onLogOut}>
        Log Out
      </button>
    </div>
  );
}

export default ManageAccount;
