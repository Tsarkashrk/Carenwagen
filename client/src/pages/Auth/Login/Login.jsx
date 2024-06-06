import { Link } from 'react-router-dom'
import Input from '../../../components/Input/Input'
import LinkButton from '../../../components/LinkButton/LinkButton'

const Login = () => {
  return (
    <div className="auth">
      <div className="auth__wrapper">
        <h1 className="auth__title">Log in</h1>
        <form className="auth__form">
          <div className="auth__email">
            <Input text="Email" />
          </div>
          <div className="auth__password">
            <Input text="Password" />
          </div>
          {/* {error && (
          <div className="auth__error">
            {Array.isArray(error) ? (
              error.map((error, index) => (
                <p key={index} className="auth__message">
                  * {error.msg}
                </p>
              ))
            ) : (
              <p className="auth__message">* {error.message}</p>
            )}
          </div>
        )} */}
          <div className="auth__processess">
            <LinkButton text="Log in" link={`/profile`} />
            <div className="auth__nav">
              Don't have an account?
              <Link className="auth__link" to="/auth/register">
                Register
              </Link>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}

export default Login
