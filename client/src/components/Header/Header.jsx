import { Link } from 'react-router-dom'

const Header = () => {
  return (
    <header className="header">
      <div className="header__wrapper">
        <Link className="header__link" to="/">
          <h1 className="header__logo">.CAREN.</h1>
          <h1 className="header__logo">WAGEN</h1>
        </Link>
        <nav className='header__navigation'>
          <ul className="header__list">
            <li className="header__item">
              <Link className="header__nav-link" to="/">
                <p className="header__nav-item">Home</p>
              </Link>
            </li>
            <li className="header__item">
              <Link className="header__nav-link" to="/catalog">
                <p className="header__nav-item">Catalog</p>
              </Link>
            </li>
            <li className="header__item">
              <Link className="header__nav-link" to="/auth/login">
                <p className="header__nav-item">Login</p>
              </Link>
            </li>
            <li className="header__item">
              <Link className="header__nav-link" to="/orders">
                <p className="header__nav-item">My orders</p>
              </Link>
            </li>
            <li className="header__item">
              <Link className="header__nav-link" to="/profile">
                <p className="header__nav-item">Profile</p>
              </Link>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  )
}

export default Header
