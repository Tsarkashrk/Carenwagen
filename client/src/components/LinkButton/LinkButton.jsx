/* eslint-disable react/prop-types */
import { Link } from 'react-router-dom'

const LinkButton = ({ text, link }) => {
  return (
    <Link className="link-button" to={link}>
      {text}
    </Link>
  )
}

export default LinkButton
