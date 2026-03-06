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
      <input
        className="account-input"
        type="email"
        placeholder="Email address"
        value={email}
        onChange={(e) => onEmailChange(e.target.value)}
      />
      <input
        className="account-input"
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => onPasswordChange(e.target.value)}
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
