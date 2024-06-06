import Actions from '../../components/Actions/Actions'
import Heading from '../../components/Heading/Heading'

const Profile = () => {
  return (
    <div className="profile">
      <div className="profile-wrapper">
        <Heading text="profile" />
        <div className="profile__content">
          <Actions />
        </div>
      </div>
    </div>
  )
}

export default Profile
