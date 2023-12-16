<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import PopUpLikeCard from '../components/PopUpLikeCard.vue';
import PostPopUp from '../components/PostPopUp/PostPopUp.vue';
import { getPictureURL } from '../functions/getPictureURL';
import PencilIcon from 'vue-material-design-icons/Pencil.vue';
import CheckIcon from 'vue-material-design-icons/Check.vue';

const initialData = () => ({
  profileuserid: null,
  profileusername: null,
  profilebio: null,
  userpicture: null,
  profilefeeling: null,
  profilefollowers: null,
  profilefollowing: null,
  profilenotfound: false,
  banned: false,
  profileposts: null,

  userid: null,
  username: null,
  bio: null,
  feeling: null,
  picture: null,

  followPopup: null,
  editUsernamePopup: false,
  editBioPopup: false,
  newPostPopup: false,
  openedPost: null,

  newPostImage: null
});

export default {
  watch: {
    $route: 'routeUpdated'
  },
  data() {
    return initialData();
  },
  async created() {
    await this.loadPage();
  },
  beforeUnmount() {
    document.body.classList.remove('no-scroll');
  },
  methods: {
    getPictureURL,
    async routeUpdated() {
      await this.loadPage();
    },
    async loadPage() {
      Object.assign(this.$data, initialData()); // restore initial state
      this.userid = parseInt(localStorage.getItem('userid'));
      this.username = localStorage.getItem('username');
      this.feeling = parseInt(localStorage.getItem('feeling'));
      this.bio = localStorage.getItem('bio');
      this.picture = parseInt(localStorage.getItem('picture'));
      document.body.classList.remove('no-scroll'); // unlock scrolling if it was locked before

      await this.getProfile();
    },
    async getProfile() {
      this.profileuserid = parseInt(this.$route.params.id);
      this.followPopup = null; // close popup (needed when changing profile from follow list)
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
              this.userpicture = 'defaultsnoopy';
              this.profilefeeling = -1;
              this.profilenotfound = true;
              break;
            case 500:
              this.profileusername = 'Server error';
              this.profilebio =
                'Sorry, we are having some problems with our servers. Please try again later.';
              this.userpicture = '';
              this.profilefeeling = -1;
              break;
          }
        });

      const response = res?.data;
      if (response) {
        this.profileusername = response.username;
        this.profilefeeling = response.feeling;
        this.profilebio = response.bio;
        this.userpicture = response.picture;
        this.profilefollowers = response.followers;
        this.profilefollowing = response.following;
        this.banned = response.banned;
        this.profileposts = response.posts;
      }
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
    setfollowPopup(type) {
      if (type == null) document.body.classList.remove('no-scroll');
      else {
        document.body.classList.add('no-scroll');
        window.scrollTo({
          top: 0,
          behavior: 'smooth'
        });
      }
      this.followPopup = type; // either 'followers' or 'following' or null
    },
    openFileDialog() {
      if (this.profileuserid != this.userid) return;
      this.$refs.fileInput.click();
    },
    async handleFileSelected(event) {
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
    async updateProfilePicture(newuserpicture) {
      const token = localStorage.getItem('token');
      const res = await this.$axios
        .put(
          `/users/${this.userid}/picture`,
          {
            newpicture: newuserpicture
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .catch((err) => {
          switch (err.response.status) {
            case 403:
              alert('Invalid credentials. Please try re-logging in.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
      if (res) {
        // get the picture id from the response
        const pictureid = res.data.pictureid;
        localStorage.setItem('picture', pictureid);
        this.picture = pictureid;
        this.userpicture = pictureid;
      }
    },
    openEditUsername() {
      document.body.classList.add('no-scroll'); // disable scrolling on page
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });
      this.editUsernamePopup = true;
    },
    closeEditUsername() {
      document.body.classList.remove('no-scroll'); // enable scrolling on page
      this.editUsernamePopup = false;
    },
    openEditBio() {
      document.body.classList.add('no-scroll'); // disable scrolling on page
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });
      this.editBioPopup = true;
    },
    closeEditBio() {
      document.body.classList.remove('no-scroll'); // enable scrolling on page
      this.editBioPopup = false;
    },
    async opennewPostPopup() {
      document.body.classList.add('no-scroll'); // disable scrolling on page
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });
      this.newPostPopup = true;
    },
    async closenewPostPopup() {
      document.body.classList.remove('no-scroll'); // enable scrolling on page
      this.newPostPopup = false;
      this.newPostImage = null;
    },
    async changeUsername() {
      this.closeEditUsername();
      const newusername = document.getElementById('username-popup-input').value;
      if (newusername === this.username) return;
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
            // no need to handle 400 because the frontend already checks the length
            case 403:
              alert('Invalid credentials. Please try re-logging in');
              break;
            case 409:
              alert('Username already taken.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    },
    async changeBio() {
      this.closeEditBio();
      const newbio = document.getElementById('bio-popup-input').value;
      if (newbio === this.bio) return;
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
            case 403:
              alert('Invalid credentials. Please try re-logging in.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    },
    async cycleFeeling() {
      if (this.profileuserid != this.userid) return;
      // feeling goes from 0 to 4, this function cycles it
      let newFeeling = (this.feeling + 1) % 5;
      const token = localStorage.getItem('token');
      this.$axios
        .put(
          `/users/${this.userid}/feeling`,
          {
            newfeeling: newFeeling
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .then(() => {
          this.feeling = newFeeling;
          this.profilefeeling = newFeeling;
          localStorage.setItem('feeling', this.feeling);
        })
        .catch((err) => {
          switch (err.response.status) {
            // no need for a 400 because the frontend already handles the feeling
            case 403:
              alert('Invalid credentials. Please try re-logging in.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    },
    async toggleFollowUser() {
      const token = localStorage.getItem('token');
      this.profilefollowers = this.profilefollowers || [];
      const alreadyFollowing = this.profilefollowers.some(
        (user) => user.userid == this.userid
      );
      if (!alreadyFollowing) {
        this.$axios
          .put(
            `/users/${this.profileuserid}/follow`,
            {},
            {
              headers: {
                Authorization: `Bearer ${token}`
              }
            }
          )
          .then(() => {
            this.profilefollowers.push({
              userid: this.userid,
              username: this.username,
              feeling: this.feeling,
              bio: this.bio,
              picture: this.picture
            });
          })
          .catch((err) => {
            switch (err.response.status) {
              case 403:
                alert('Invalid credentials. Please try re-logging in.');
                break;
              case 404:
                alert('User not found.');
                break;
              case 500:
                alert('Server error. Please try again later.');
                break;
            }
          });
      } else {
        this.$axios
          .delete(`/users/${this.profileuserid}/follow`, {
            headers: {
              Authorization: `Bearer ${token}`
            },
            data: {}
          })
          .then(() => {
            this.profilefollowers = this.profilefollowers.filter(
              (user) => user.userid != this.userid
            );
          })
          .catch((err) => {
            switch (err.response.status) {
              case 403:
                alert('Invalid credentials. Please try re-logging in.');
                break;
              case 404:
                alert('User not found.');
                break;
              case 500:
                alert('Server error. Please try again later.');
                break;
            }
          });
      }
    },
    toggleBanUser() {
      const token = localStorage.getItem('token');
      if (!this.banned) {
        this.$axios
          .put(
            `/users/${this.profileuserid}/ban`,
            {},
            {
              headers: {
                Authorization: `Bearer ${token}`
              }
            }
          )
          .then(() => {
            this.banned = true;
          })
          .catch((err) => {
            switch (err.response.status) {
              case 500:
                alert('Server error. Please try again later.');
                break;
            }
          });
      } else {
        this.$axios
          .delete(`/users/${this.profileuserid}/ban`, {
            headers: {
              Authorization: `Bearer ${token}`
            },
            data: {}
          })
          .then(() => {
            this.banned = false;
          })
          .catch((err) => {
            switch (err.response.status) {
              case 500:
                alert('Server error. Please try again later.');
                break;
            }
          });
      }
    },
    openNewPostFileDialog() {
      this.$refs.newPostFileInput.click();
    },
    async handlePostFileSelected() {
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
          vuethis.newPostImage = base64image;
        };
        img.src = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    async publishNewPost() {
      const caption = document.getElementById(
        'profile-page-newpost-popup-input'
      ).value;
      if (this.newPostImage == null) {
        alert('Please select an image.');
        return;
      }
      const token = localStorage.getItem('token');
      this.$axios
        .put(
          `/posts`,
          {
            caption: caption,
            image: this.newPostImage
          },
          {
            headers: {
              Authorization: `Bearer ${token}`
            }
          }
        )
        .then((res) => {
          this.closenewPostPopup();
          document.body.classList.remove('no-scroll');
          this.newPostImage = null;
          let newPost = {
            postid: res.data.postid,
            caption: caption,
            picture: res.data.picture,
            createdat: res.data.createdat,
            likecount: res.data.likecount
          };

          this.profileposts = [newPost, ...(this.profileposts ?? [])];
        })
        .catch((err) => {
          console.log(err);
          switch (err.response.status) {
            case 403:
              alert('Invalid credentials. Please try re-logging in.');
              break;
            case 500:
              alert('Server error. Please try again later.');
              break;
          }
        });
    },
    openPost(post) {
      this.openedPost = post;
    },
    closePost() {
      this.openedPost = null;
    }
  },
  components: {
    HeaderComponent,
    PopUpLikeCard,
    PostPopUp,
    PencilIcon,
    CheckIcon
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
            :src="getPictureURL(this.userpicture)"
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
              <button @click="cycleFeeling()" class="profile-info-feeling">
                <span v-if="this.profilefeeling === 0">😐</span>
                <span v-if="this.profilefeeling === 1">😀</span>
                <span v-if="this.profilefeeling === 2">😍</span>
                <span v-if="this.profilefeeling === 3">😡</span>
                <span v-if="this.profilefeeling === 4">😭</span>
              </button>
              <div class="profile-info-userame">{{ profileusername }}</div>
              <PencilIcon
                v-if="this.profileuserid == this.userid"
                @click="this.openEditUsername()"
                class="edit-icon"
                :size="24"
              />
              <div class="posts-count">
                {{ this.profileposts?.length || 0 }} Posts
              </div>
              <button
                v-if="
                  this.profileuserid != this.userid && !this.profilenotfound
                "
                class="action-button"
                @click="this.toggleFollowUser()"
              >
                {{
                  this.profilefollowers
                    ? this.profilefollowers.some(
                        (user) => user.userid == this.userid
                      )
                      ? 'Unfollow'
                      : 'Follow'
                    : 'Follow'
                }}
              </button>
              <button
                v-if="
                  !this.profilenotfound && this.profileuserid != this.userid
                "
                class="action-button"
                @click="toggleBanUser()"
              >
                {{ this.banned ? 'Unban' : 'Ban' }}
              </button>
            </div>
            <div class="profile-info-bio-container">
              <div class="profile-info-bio">
                {{ profilebio || 'No bio, yet.' }}

                <PencilIcon
                  v-if="this.profileuserid == this.userid"
                  @click="this.openEditBio()"
                  class="edit-icon"
                  :size="20"
                />
              </div>
            </div>
          </div>
        </div>
        <div class="profile-info-social">
          <button @click="setfollowPopup('followers')">
            Followers
            <div class="profile-info-social-number">
              {{ getFollowNumber(this.profilefollowers) }}
            </div>
          </button>
          <button @click="setfollowPopup('following')">
            Following
            <div class="profile-info-social-number">
              {{ getFollowNumber(this.profilefollowing) }}
            </div>
          </button>
        </div>
      </div>
      <div class="profile-page-posts">
        <button
          @click="this.opennewPostPopup()"
          v-if="this.profileuserid == this.userid"
          class="profile-page-newpost-card"
        >
          <div class="profile-page-newpost-card-plus">+</div>
          <div class="profile-page-newpost-card-title">New Post</div>
        </button>
        <button
          v-if="this.profileposts"
          v-for="post in this.profileposts"
          @click="openPost(post)"
          :key="post.postid"
          class="profile-page-post-card"
        >
          <img :src="getPictureURL(post.picture)" alt="" />
          <div class="profile-page-post-card-text">
            {{ post.caption || 'No caption' }}
          </div>
        </button>
      </div>
    </div>
  </div>
  <div
    @click="setfollowPopup(null)"
    v-if="this.followPopup != null"
    class="profile-page-follow-popup-outer-container"
  >
    <div @click.stop="() => {}" class="profile-page-follow-popup-container">
      <div
        v-if="
          (followPopup == 'followers'
            ? this.profilefollowers
            : this.profilefollowing) == null
        "
      >
        No users yet! :(
      </div>
      <PopUpLikeCard
        v-for="user in followPopup == 'followers'
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
  <div
    @click="this.closenewPostPopup()"
    v-if="this.newPostPopup"
    class="profile-page-newpost-popup-outer-container"
  >
    <div @click.stop="() => {}" class="profile-page-newpost-popup-container">
      <button
        v-if="this.newPostImage == null"
        @click="openNewPostFileDialog()"
        class="profile-page-newpost-popup-image"
      >
        Select an image
      </button>
      <input
        type="file"
        ref="newPostFileInput"
        @change="handlePostFileSelected"
        style="display: none"
      />
      <!-- data:image/png;base64, prefix needed because we strip it when going to the server -->
      <img
        v-if="this.newPostImage"
        :src="'data:image/png;base64,' + this.newPostImage"
        class="profile-page-newpost-popup-image"
      />
      <textarea
        id="profile-page-newpost-popup-input"
        placeholder="Write your caption here..."
      />
      <button
        @click="publishNewPost()"
        class="profile-page-newpost-popup-publish-button"
      >
        Publish Post
      </button>
    </div>
  </div>
  <PostPopUp
    v-if="this.openedPost"
    :postid="this.openedPost.postid"
    :userid="this.profileuserid"
    :name="this.profileusername"
    :feeling="this.profilefeeling"
    :picture="this.openedPost.picture"
    :userPicture="this.userpicture"
    :likeCount="this.openedPost.likecount"
    :date="this.openedPost.createdat"
    :caption="this.openedPost.caption"
    :closePost="this.closePost"
  />
</template>

<style lang="scss" scoped>
body.no-scroll {
  overflow: hidden;
}
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
        @media screen and (max-width: 1000px) {
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
            margin-right: 10px;
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
            .profile-info-feeling {
              font-size: 30px;
              border: none;
              background-color: transparent;
              margin-right: 3px;
              transition: all 0.3s ease-in-out;
              &:hover {
                transform: scale(1.1);
              }
            }
            .posts-count {
              border: 2px solid orange;
              border-radius: 8px;
              background-color: transparent;
              color: rgb(78, 78, 78);
              font-size: 15px;
              font-weight: bold;
              margin-right: 8px;
              padding: 1px 5px;
              transition: all 0.2s ease-in-out;
              color: orange;
            }
            .action-button {
              border: 2px solid rgb(78, 78, 78);
              border-radius: 8px;
              background-color: transparent;
              color: rgb(78, 78, 78);
              font-size: 15px;
              font-weight: bold;
              margin-right: 8px;
              transition: all 0.2s ease-in-out;
              &:hover {
                cursor: pointer;
                transform: scale(1.05);
              }
            }
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
      .profile-page-post-card {
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
        justify-content: space-between;
        transition: all 0.2s ease-in-out;
        padding: 10px;
        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
          border-radius: 12px;
          margin-bottom: 10px;
        }
        .profile-page-post-card-text {
          width: 100%;
          max-height: 58px;
          padding: 10px;
          font-family: 'Courier New', Courier, monospace;
          font-size: 16px;
          overflow: hidden;
          text-overflow: ellipsis;
          word-wrap: break-word;
        }
        &:hover {
          transform: scale(1.01);
          cursor: pointer;
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
.profile-page-newpost-popup-outer-container {
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
  .profile-page-newpost-popup-container {
    padding: 15px 15px;
    background-color: white;
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: space-between;
    width: 400px;
    height: 70vh;
    font-family: 'Courier New', Courier, monospace;
    .profile-page-newpost-popup-image {
      width: 100%;
      max-height: 60%;
      border-radius: 8px;
      margin-right: 8px;
    }
    img {
      min-height: 200px;
      max-height: 60%;
      width: auto;
      object-fit: contain;
    }
    textarea {
      margin: 10px 0px;
      width: 100%;
      height: 100%;
      border: none;
      border-radius: 8px;
      margin-right: 8px;
      font-weight: bold;
      font-size: 18px;
      resize: none;
      outline: none;
    }

    .profile-page-newpost-popup-publish-button {
      border-radius: 8px;
      border: 2px solid rgb(71, 71, 71);
      color: rgb(71, 71, 71);
      padding: 5px;
      transition: all 0.2s ease-in-out;
      background-color: transparent;
      font-weight: bold;
      font-size: 18px;
      &:hover {
        color: white;
        background-color: orange;
        border: 2px solid orange;
        transform: scale(1.05);
        cursor: pointer;
      }
    }
  }
}
</style>
