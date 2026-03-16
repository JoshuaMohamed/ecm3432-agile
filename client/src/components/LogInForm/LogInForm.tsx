import AccountFormFields from "../AccountFormFields/AccountFormFields";
import "../AccountForm.css";

interface LogInFormProps {
  email: string;
  password: string;
  onEmailChange: (value: string) => void;
  onPasswordChange: (value: string) => void;
  onSubmit: () => void;
  onSwitchToSignUp: () => void;
}

function LogInForm({
  email,
  password,
  onEmailChange,
  onPasswordChange,
  onSubmit,
  onSwitchToSignUp,
}: LogInFormProps) {
  return (
    <div className="account-form">
      <h2 className="account-heading">Log In</h2>
      <AccountFormFields
        email={email}
        password={password}
        onEmailChange={onEmailChange}
        onPasswordChange={onPasswordChange}
      />
      <button className="account-button" onClick={onSubmit}>
        Log In
      </button>
      <p className="account-switch">
        Don't have an account?{" "}
        <span className="account-link" onClick={onSwitchToSignUp}>
          Sign Up
        </span>
      </p>
    </div>
  );
}

export default LogInForm;
