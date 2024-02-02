import logo from '../../assets/logo.svg';

import './Header.css';

function Header() {
  return (
    <div class="HeaderContainer">
        <img src={logo} class="HeaderLogo" alt="logo" />
        <h1 class="Title">Bobbitt</h1>
    </div>
  )
}

export default Header;