import "../AccountForm.css";

type AccountFormFieldsProps = {
  email: string;
  password: string;
  onEmailChange: (value: string) => void;
  onPasswordChange: (value: string) => void;
};

function AccountFormFields({
  email,
  password,
  onEmailChange,
  onPasswordChange,
}: AccountFormFieldsProps) {
  return (
    <>
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
    </>
  );
}

export default AccountFormFields;
