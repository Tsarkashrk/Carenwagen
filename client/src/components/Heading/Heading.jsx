/* eslint-disable react/prop-types */
const Heading = ({ text }) => {
  const upperCaseText = text.toUpperCase()

  return <h2 className="heading">{upperCaseText}</h2>
}

export default Heading
