import AccountFormFields from "../AccountFormFields/AccountFormFields";
import "../AccountForm.css";

interface SignUpFormProps {
  email: string;
  password: string;
  onEmailChange: (value: string) => void;
  onPasswordChange: (value: string) => void;
  onSubmit: () => void;
  onSwitchToLogin: () => void;
}

function SignUpForm({
  email,
  password,
  onEmailChange,
  onPasswordChange,
  onSubmit,
  onSwitchToLogin,
}: SignUpFormProps) {
  return (
    <div className="account-form">
      <h2 className="account-heading">Sign Up</h2>
      <AccountFormFields
        email={email}
        password={password}
        onEmailChange={onEmailChange}
        onPasswordChange={onPasswordChange}
      />
      <button className="account-button" onClick={onSubmit}>
        Sign Up
      </button>
      <p className="account-switch">
        Already have an account?{" "}
        <span className="account-link" onClick={onSwitchToLogin}>
          Log In
        </span>
      </p>
    </div>
  );
}

export default SignUpForm;
