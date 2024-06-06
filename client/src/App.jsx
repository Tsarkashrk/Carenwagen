import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout/Layout'
import Home from './pages/Home/Home'
import Profile from './pages/Profile/Profile'
import Catalog from './pages/Catalog/Catalog'
import Login from './pages/Auth/Login/Login'
import Register from './pages/Auth/Register/Register'
import ProductDetails from './components/ProductDetails/ProductDetails'
import Orders from './pages/Orders/Orders'

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="profile" element={<Profile />} />
          <Route path="car/:id/details" element={<ProductDetails />} />
          <Route path="catalog" element={<Catalog />} />
          <Route path="orders" element={<Orders />} />
          <Route path="auth/login" element={<Login />} />
          <Route path="auth/register" element={<Register />} />
        </Route>
      </Routes>
    </Router>
  )
}

export default App
