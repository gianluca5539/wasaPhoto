<script>
import UserHeaderCard from './UserHeaderCard.vue';
import SearchIcon from 'vue-material-design-icons/Magnify.vue';
import CloseIcon from 'vue-material-design-icons/Close.vue';

export default {
  name: 'HeaderComponent',
  data() {
    return {
      token: null,
      searchOpen: false
    };
  },
  props: {
    userid: {
      type: Number,
      required: true
    },
    username: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    picture: {
      type: Number,
      required: true
    }
  },
  methods: {
    goToHomePage() {
      this.$router.push('/');
    },
    toggleSearchOpen() {
      this.searchOpen = !this.searchOpen;
      setTimeout(() => {
        if (this.searchOpen) {
          document.getElementById('header-user-search-input').focus();
        }
      }, 100);
    },
    searchUser() {
      const searchInput = document.getElementById('header-user-search-input');
      if (searchInput.value.length > 0) {
        const username = searchInput.value;
        // get userid from username
        this.$axios
          .get(`/userids/${username}`, {
            headers: {
              Authorization: `Bearer ${this.token}`
            }
          })
          .then((response) => {
            const userid = response.data.id;
            this.toggleSearchOpen();
            this.$router.push(`/profile/${userid}`);
          })
          .catch((error) => {
            switch (error.response.status) {
              case 404:
                this.toggleSearchOpen();
                this.$router.push(`/profile/0`);
                break;
              case 500:
                alert('We are having some problems, try again later.');
                break;
            }
          });
      }
    }
  },
  components: { UserHeaderCard, SearchIcon, CloseIcon },
  created() {
    this.token = localStorage.getItem('token');
  }
};
</script>

<template>
  <div class="homepage-header-container">
    <div class="logo-search-container">
      <button @click="this.toggleSearchOpen()" class="header-search-button">
        <SearchIcon v-if="!this.searchOpen" :size="40" />
      </button>
      <button
        v-if="!this.searchOpen"
        @click="this.goToHomePage()"
        class="header-logo-button"
      >
        WASAPhoto
      </button>
      <input
        v-if="this.searchOpen"
        type="text"
        placeholder="Search users..."
        id="header-user-search-input"
        @keyup.enter="this.searchUser()"
      />
      <button
        v-if="this.searchOpen"
        @click="this.toggleSearchOpen()"
        class="header-search-button"
      >
        <CloseIcon :size="45" />
      </button>
    </div>
    <UserHeaderCard
      :userid="userid"
      :username="username"
      :feeling="feeling"
      :picture="picture"
    />
  </div>
</template>

<style lang="scss">
.homepage-header-container {
  width: 100%;
  max-width: 100%;
  height: 70px;
  font-family: 'Courier New', Courier, monospace;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 30px;
  position: absolute;
  top: 0px;
  right: 0px;
  .logo-search-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 200px;
    button {
      font-size: 45px;
      font-weight: bold;
      color: white;
      background: transparent;
      border: none;
      outline: none;
      cursor: pointer;
      transition: all 0.3s ease;
      &:hover {
        color: orange;
        transform: scale(1.01);
      }

      &:first-child {
        margin-right: 8px;
        padding-bottom: 4px;
      }
    }
    input {
      width: 200px;
      height: 40px;
      font-size: 20px;
      font-weight: bold;
      color: white;
      background: transparent;
      border: none;
      outline: none;
      border-bottom: 2px solid white;
      transition: all 0.3s ease;
      &:focus {
        border-bottom: 2px solid orange;
      }
    }
  }
}
</style>
