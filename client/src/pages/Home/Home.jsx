import Heading from '../../components/Heading/Heading.jsx'
import ProductList from '../../components/ProductList/ProductList.jsx'
import CARS_DATA from '../../constants/carsData.js'

const Home = () => {
  return (
    <div className="home">
      <div className="home__wrapper">
        <div className="home__preview">
          <h1 className="home__preview-text">FIND YOUR CAR</h1>
        </div>
        <Heading text="Last updates" />
        <div className="home__products">
          <ProductList carsData={CARS_DATA} />
        </div>
      </div>
    </div>
  )
}

export default Home
