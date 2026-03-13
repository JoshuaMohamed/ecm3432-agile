import { useEffect, useState } from "react";

import "./AccountPage.css";
import TopBar from "../../components/TopBar/TopBar";
import SignUpForm from "../../components/SignUpForm/SignUpForm";
import LogInForm from "../../components/LogInForm/LogInForm";
import ManageAccount from "../../components/ManageAccount/ManageAccount";

type View = "signup" | "login" | "manage";

function AccountPage() {
  const [hasSession, setHasSession] = useState<boolean>(false);
  const [view, setView] = useState<View>("signup");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loggedInEmail, setLoggedInEmail] = useState<string | null>(null);

  async function validateSession() {
    try {
      const res = await fetch("http://localhost:8080/session", {
        method: "GET",
        credentials: "include",
      });

      if (!res.ok) {
        setHasSession(false);
        setLoggedInEmail(null);
        setView("signup");
        return;
      }

      const body = await res.json();
      const sessionEmail = body?.Data?.email;

      setHasSession(true);
      setLoggedInEmail(typeof sessionEmail === "string" ? sessionEmail : null);
      setView("manage");
    } catch {
      setHasSession(false);
      setLoggedInEmail(null);
      setView("signup");
    }
  }

  useEffect(() => {
    validateSession();
  }, []);

  async function handleSignUp() {
    setError(null);
    try {
      const res = await fetch("http://localhost:8080/signup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ email, password, role: "tourist" }),
      });
      if (!res.ok) {
        const body = await res.json();
        setError(body.Message || "Failed to create account");
        return;
      }
      setLoggedInEmail(email);
      setEmail("");
      setPassword("");
      await validateSession();
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }
  }

  async function handleLogIn() {
    setError(null);
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ email, password }),
      });
      if (!res.ok) {
        const body = await res.json();
        setError(body.Message || "Invalid email or password");
        return;
      }
      setLoggedInEmail(email);
      setEmail("");
      setPassword("");
      await validateSession();
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }
  }

  async function handleLogOut() {
    setError(null);

    try {
      await fetch("http://localhost:8080/logout", {
        method: "DELETE",
        credentials: "include",
      });
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }

    setHasSession(false);
    setLoggedInEmail(null);
    setEmail("");
    setPassword("");
    setView("signup");
  }

  function switchView(target: View) {
    setError(null);
    setView(target);
  }

  return (
    <div className="page">
      <TopBar />

      <main className="content">
        {!hasSession && view === "signup" && (
          <SignUpForm
            email={email}
            password={password}
            onEmailChange={setEmail}
            onPasswordChange={setPassword}
            onSubmit={handleSignUp}
            onSwitchToLogin={() => switchView("login")}
          />
        )}

        {!hasSession && view === "login" && (
          <LogInForm
            email={email}
            password={password}
            onEmailChange={setEmail}
            onPasswordChange={setPassword}
            onSubmit={handleLogIn}
            onSwitchToSignUp={() => switchView("signup")}
          />
        )}

        {hasSession && (
          <ManageAccount email={loggedInEmail} onLogOut={handleLogOut} />
        )}

        {error && <p className="error">{error}</p>}
      </main>
    </div>
  );
}

export default AccountPage;
