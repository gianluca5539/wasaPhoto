<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import PopUpLikeCard from '../components/PopUpLikeCard.vue';
import { getPictureURL } from '../functions/getPictureURL';
import PencilIcon from 'vue-material-design-icons/Pencil.vue';
import CheckIcon from 'vue-material-design-icons/Check.vue';

export default {
  watch: {
    $route: 'getProfile'
  },
  data() {
    return {
      profileuserid: null,
      profileusername: null,
      profilebio: null,
      profilepicture: null,
      profilefeeling: null,
      profilefollowers: null,
      profilefollowing: null,

      userid: null,
      username: null,
      bio: null,
      feeling: null,
      picture: null,

      followpopup: null,
      editUsernamePopup: false,
      editBioPopup: false
    };
  },
  methods: {
    getPictureURL,
    async getProfile() {
      this.followpopup = null; // close popup (needed when changing profile from follow list)
      const token = localStorage.getItem('token');
      const res = await this.$axios
        .get(`/users/${this.profileuserid}/profile`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        .catch((err) => {
          switch (err.response.status) {
            case 404:
              this.profileusername = 'User not found';
              this.profilebio =
                'Sorry, we looked everywhere but we could not find this user.';
              this.profilepicture = '';
              this.profilefeeling = -1;
              break;
            case 500:
              this.profileusername = 'Server error';
              this.profilebio =
                'Sorry, we are having some problems with our servers. Please try again later.';
              this.profilepicture = '';
              this.profilefeeling = -1;
              break;
            default:
              this.$router.push('/serviceunavailable');
              break;
          }
        });

      const response = res.data;
      this.profileusername = response.username;
      this.profilefeeling = response.feeling;
      this.profilebio = response.bio;
      this.profilepicture = response.picture;
      this.profilefollowers = response.followers;
      this.profilefollowing = response.following;
    },
    getFollowNumber(x) {
      if (x === null) return 0;
      let length = x.length;
      if (length > 1000000) {
        return `${(length / 1000000).toFixed(1)}M`;
      } else if (length > 10000) {
        return `${(length / 1000).toFixed(1)}K`;
      } else {
        return length;
      }
    },
    setFollowPopup(type) {
      this.followpopup = type; // either 'followers' or 'following' or null
    },
    openFileDialog() {
      this.$refs.fileInput.click();
    },
    handleFileSelected(event) {
      const file = event.target.files[0];
      if (!file.type.startsWith('image/')) {
        alert('Please select an image file.');
        event.target.value = '';
        return;
      }
      let vuethis = this; // needed to access this inside the onload function
      // get the base64 string of the image in PNG
      const reader = new FileReader();
      reader.onload = (e) => {
        const img = new Image();
        img.onload = function () {
          const canvas = document.createElement('canvas');
          canvas.width = this.naturalWidth;
          canvas.height = this.naturalHeight;

          const ctx = canvas.getContext('2d');
          ctx.drawImage(this, 0, 0);

          const pngBase64String = canvas.toDataURL('image/png');
          const base64image = pngBase64String.split(',')[1];
          vuethis.updateProfilePicture(base64image);
        };
        img.src = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    async updateProfilePicture(newprofilepicture) {
      const token = localStorage.getItem('token');
      const res = await this.$axios
        .put(
          `/users/${this.userid}/picture`,
          {
            newpicture: newprofilepicture
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .catch((err) => {
          alert('Error updating profile picture. Please try again later.');
        });
      if (res) {
        // get the picture id from the response
        const pictureid = res.data.pictureid;
        localStorage.setItem('picture', pictureid);
        this.picture = pictureid;
        this.profilepicture = pictureid;
      }
    },
    openEditUsername() {
      this.editUsernamePopup = true;
    },
    closeEditUsername() {
      this.editUsernamePopup = false;
    },
    openEditBio() {
      this.editBioPopup = true;
    },
    closeEditBio() {
      this.editBioPopup = false;
    },
    changeUsername() {
      this.closeEditUsername();
      const newusername = document.getElementById('username-popup-input').value;
      if (newusername.length < 3 || newusername.length > 16) {
        alert('Username must be between 3 and 16 characters long.');
        return;
      }
      const token = localStorage.getItem('token');
      this.$axios
        .put(
          `/users/${this.userid}/username`,
          {
            newusername: newusername
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .then((res) => {
          localStorage.setItem('username', newusername);
          this.username = newusername;
          this.profileusername = newusername;
        })
        .catch((err) => {
          switch (err.response.status) {
            case 400:
              alert('Username already taken.');
              break;
            case 401:
              alert('Invalid token.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    },
    changeBio() {
      this.closeEditBio();
      const newbio = document.getElementById('bio-popup-input').value;
      const token = localStorage.getItem('token');
      this.$axios
        .put(
          `/users/${this.userid}/bio`,
          {
            newbio: newbio
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .then((res) => {
          localStorage.setItem('bio', newbio);
          this.bio = newbio;
          this.profilebio = newbio;
        })
        .catch((err) => {
          switch (err.response.status) {
            case 401:
              alert('Invalid token.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    }
  },
  components: { HeaderComponent, PopUpLikeCard, PencilIcon, CheckIcon },
  async created() {
    this.userid = parseInt(localStorage.getItem('userid'));
    this.username = localStorage.getItem('username');
    this.feeling = parseInt(localStorage.getItem('feeling'));
    this.bio = localStorage.getItem('bio');
    this.picture = localStorage.getItem('picture');
    this.profileuserid = this.$route.params.id;

    await this.getProfile();
  }
};
</script>

<template>
  <div class="profile-page-container">
    <HeaderComponent
      :userid="this.userid"
      :username="this.username"
      :feeling="this.feeling"
      :picture="this.picture"
    />
    <div class="profile-page-content">
      <div class="profile-info-container">
        <div class="profile-info-data-container">
          <img
            @click="openFileDialog()"
            :src="getPictureURL(profilepicture)"
            alt=""
          />
          <input
            type="file"
            ref="fileInput"
            @change="handleFileSelected"
            style="display: none"
          />
          <div class="profile-info-data">
            <div class="profile-info-username-container">
              <div class="profile-info-userame">{{ profileusername }}</div>
              <PencilIcon
                @click="this.openEditUsername()"
                class="edit-icon"
                :size="24"
              />
            </div>
            <div class="profile-info-bio-container">
              <div class="profile-info-bio">
                {{ profilebio || 'No bio, yet.' }}
              </div>
              <PencilIcon
                @click="this.openEditBio()"
                class="edit-icon"
                :size="20"
              />
            </div>
          </div>
        </div>
        <div class="profile-info-social">
          <button @click="setFollowPopup('followers')">
            Followers
            <div class="profile-info-social-number">
              {{ getFollowNumber(this.profilefollowers) }}
            </div>
          </button>
          <button @click="setFollowPopup('following')">
            Following
            <div class="profile-info-social-number">
              {{ getFollowNumber(this.profilefollowing) }}
            </div>
          </button>
        </div>
      </div>
      <div class="profile-page-posts">
        <button
          v-if="this.profileuserid == this.userid"
          class="profile-page-newpost-card"
        >
          <div class="profile-page-newpost-card-plus">+</div>
          <div class="profile-page-newpost-card-title">New Post</div>
        </button>
      </div>
    </div>
  </div>
  <div
    @click="setFollowPopup(null)"
    v-if="this.followpopup != null"
    class="profile-page-follow-popup-outer-container"
  >
    <div @click.stop="() => {}" class="profile-page-follow-popup-container">
      <div
        v-if="
          (followpopup == 'followers'
            ? this.profilefollowers
            : this.profilefollowing) == null
        "
      >
        No users yet! :(
      </div>
      <PopUpLikeCard
        v-for="user in followpopup == 'followers'
          ? this.profilefollowers
          : this.profilefollowing"
        :key="user.userid"
        :userid="user.userid"
        :name="user.username"
        :feeling="user.feeling"
        :bio="user.bio"
        :picture="user.picture"
      />
    </div>
  </div>
  <div
    @click="this.closeEditUsername()"
    v-if="this.editUsernamePopup"
    class="profile-page-username-popup-outer-container"
  >
    <div @click.stop="() => {}" class="profile-page-username-popup-container">
      <input
        id="username-popup-input"
        minlength="3"
        maxlength="16"
        type="text"
        :value="this.username"
      />
      <CheckIcon @click="this.changeUsername()" class="done-icon" :size="28" />
    </div>
  </div>
  <div
    @click="this.closeEditBio()"
    v-if="this.editBioPopup"
    class="profile-page-bio-popup-outer-container"
  >
    <div @click.stop="() => {}" class="profile-page-bio-popup-container">
      <textarea id="bio-popup-input" :value="this.bio" />
      <CheckIcon @click="this.changeBio()" class="done-icon" :size="28" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.profile-page-container {
  width: 100vw;
  min-height: 100vh;
  height: fit-content;
  display: flex;
  flex-direction: column;
  background: url('../assets/images/homepage-background.jpg') no-repeat fixed;
  .profile-page-content {
    margin-top: 80px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0px 20px;
    .profile-info-container {
      width: 80%;
      padding: 18px;
      display: flex;
      flex-direction: row;
      align-self: flex-start;
      align-items: center;
      justify-content: space-between;
      background-color: white;
      border-radius: 12px;
      margin-bottom: 50px;
      box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
      @media screen and (max-width: 800px) {
        flex-direction: column;
        justify-content: center;
      }
      .profile-info-data-container {
        display: flex;
        flex-direction: row;
        align-self: flex-start;
        align-items: center;
        justify-content: flex-start;
        @media screen and (max-width: 800px) {
          flex-direction: column;
          align-self: center;
        }
        img {
          width: 150px;
          height: 150px;
          min-height: 150px;
          min-width: 150px;
          border-radius: 10%;
          margin-right: 20px;
          cursor: pointer;
          transition: all 0.2s ease-in-out;
          &:hover {
            transform: scale(1.05);
          }
        }
        .profile-info-data {
          display: flex;
          height: 100%;
          flex-direction: column;
          align-items: flex-start;
          justify-content: flex-start;
          text-overflow: ellipsis;
          overflow: hidden;
          word-wrap: break-word;
          @media screen and (max-width: 800px) {
            align-items: center;
            text-align: center;
          }
          .edit-icon {
            transition: all 0.3s ease-in-out;
            &:hover {
              transform: scale(1.05);
              color: orange;
              cursor: pointer;
            }
          }
          .profile-info-username-container {
            display: flex;
            flex-direction: row;
            align-items: center;
            .profile-info-userame {
              font-size: 24px;
              font-weight: bold;
              max-lines: 1;
              margin-right: 5px;
            }
          }
          .profile-info-bio-container {
            display: flex;
            flex-direction: row;
            align-items: center;
            .profile-info-bio {
              font-size: 18px;
              font-weight: 300;
              max-lines: 3;
              margin-right: 5px;
            }
          }
        }
      }
      .profile-info-social {
        font-family: 'Courier New', Courier, monospace;
        font-size: 18px;
        font-weight: bold;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        button {
          border: none;
          background-color: transparent;
          color: rgb(78, 78, 78);
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          margin-bottom: 10px;
          transition: all 0.2s ease-in-out;
          &:hover {
            cursor: pointer;
            transform: scale(1.05);
          }
          .profile-info-social-number {
            font-size: 24px;
            font-weight: bold;
            color: black;
            border-radius: 10px;
            width: 100%;
            background-color: rgb(235, 235, 235);
          }
        }
      }
    }
    .profile-page-posts {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      width: 100%;
      padding: 0px 30px;
      align-items: center;
      @media screen and (max-width: 700px) {
        flex-direction: column;
        align-items: center;
      }
      .profile-page-newpost-card {
        width: 300px;
        height: 400px;
        margin-right: 30px;
        margin-bottom: 30px;
        border-radius: 12px;
        border: none;
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
        background-color: white;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
        transition: all 0.2s ease-in-out;
        &:hover {
          transform: scale(1.01);
        }
        .profile-page-newpost-card-plus {
          border-radius: 50%;
          background-color: rgb(248, 248, 248);
          color: gray;
          width: 220px;
          height: 220px;
          font-size: 100px;
          display: flex;
          align-items: center;
          justify-content: center;
          text-align: center;
          padding-bottom: 20px;
        }
        .profile-page-newpost-card-title {
          color: gray;
          font-size: 26px;
          font-family: 'Courier New', Courier, monospace;
          font-weight: bold;
        }
      }
    }
  }
}
.profile-page-follow-popup-outer-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 3;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  .profile-page-follow-popup-container {
    width: 450px;
    height: 75%;
    padding: 20px;
    background-color: white;
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow-y: scroll;
  }
}
.profile-page-username-popup-outer-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 3;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  .profile-page-username-popup-container {
    padding: 12px 12px;
    background-color: white;
    border-radius: 12px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 400px;
    height: 70px;
    input {
      height: 100%;
      width: 100%;
      border: none;
      border-radius: 8px;
      margin-right: 8px;
      font-family: 'Courier New', Courier, monospace;
      font-weight: bold;
      font-size: 22px;
      outline: none;
    }

    .done-icon {
      border-radius: 50%;
      border: 2px solid rgb(71, 71, 71);
      padding: 5px;
      transition: all 0.2s ease-in-out;
      &:hover {
        color: white;
        background-color: orange;
        border: 2px solid orange;
        transform: scale(1.1);
        cursor: pointer;
      }
    }
  }
}
.profile-page-bio-popup-outer-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 3;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  .profile-page-bio-popup-container {
    padding: 12px 12px;
    background-color: white;
    border-radius: 12px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 400px;
    min-height: 70px;
    textarea {
      width: 100%;
      border: none;
      border-radius: 8px;
      margin-right: 8px;
      font-family: 'Courier New', Courier, monospace;
      font-weight: bold;
      font-size: 18px;
      resize: none;
      outline: none;
    }

    .done-icon {
      border-radius: 50%;
      border: 2px solid rgb(71, 71, 71);
      padding: 5px;
      transition: all 0.2s ease-in-out;
      &:hover {
        color: white;
        background-color: orange;
        border: 2px solid orange;
        transform: scale(1.1);
        cursor: pointer;
      }
    }
  }
}
</style>
