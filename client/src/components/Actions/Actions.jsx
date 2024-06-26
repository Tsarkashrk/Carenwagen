import { Link } from 'react-router-dom'
import { useState } from 'react'
import ACTIONS from '../../constants/actions'

import AddCar from '../AddCar/AddCar'
import UpdateCar from '../UpdateCar/UpdateCar'
import DeleteCar from '../DeleteCar/DeleteCar'
import Orders from '../Orders/Orders'
import DeleteOrder from '../DeleteOrder/DeleteOrder'
import CreateOrder from '../CreateOrder/CreateOrder'
import EditOrder from '../EditOrder/EditOrder'
import GetCars from '../GetCars/GetCars'

const Actions = () => {
  const [actionId, setActionId] = useState(1)

  let result

  switch (actionId) {
    case 1:
      result = <AddCar />
      break
    case 2:
      result = <UpdateCar />
      break
    case 3:
      result = <DeleteCar />
      break
    case 4:
      result = <CreateOrder />
      break
    case 5:
      result = <Orders />
      break
    case 6:
      result = <DeleteOrder />
      break
    case 7:
      result = <EditOrder />
      break
    case 8:
      result = <GetCars />
      break
  }

  return (
    <div className="actions">
      <div className="actions__wrapper">
        <nav className="actions__nav">
          <ul className="actions__list">
            {ACTIONS.map((action) => (
              <li key={action.id} className="actions__item" onClick={() => setActionId(action.id)}>
                {action.icon}
                <Link className="actions__link">{action.text}</Link>
              </li>
            ))}
          </ul>
        </nav>
        {result}
      </div>
    </div>
  )
}

export default Actions
