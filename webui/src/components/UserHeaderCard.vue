<script>
import { logout } from '../functions/auth/logout';
export default {
  name: 'UserHeaderCard',
  data() {
    return {
      optionsOpen: () => window.innerWidth > 1000
    };
  },
  methods: {
    async toggleOptions() {
      if (window.innerWidth <= 1000) {
        this.optionsOpen = !this.optionsOpen;
      }
    },
    checkOptionsOpen() {
      this.optionsOpen = window.innerWidth > 1000;
    },
    openMyProfile() {
      let userid = localStorage.getItem('userid');
      this.$router.push(`/profile/${userid}`);
    },
    logout
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
    pictureURL: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    }
  },
  created() {
    window.addEventListener('resize', this.checkOptionsOpen);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.checkOptionsOpen());
  }
};
</script>

<template>
  <button
    v-if="!optionsOpen"
    class="user_header_card-container"
    @click="toggleOptions"
  >
    <div class="user_header_card-picture">
      <img :src="pictureURL" alt="" />
      <div class="user_header_card-feeling">
        <div>
          <span v-if="feeling === 0">😐</span>
          <span v-if="feeling === 1">😀</span>
          <span v-if="feeling === 2">😍</span>
          <span v-if="feeling === 3">😡</span>
          <span v-if="feeling === 4">😭</span>
        </div>
      </div>
    </div>
    <div class="user_header_card-name">{{ username }}</div>
  </button>
  <div class="user_header_card-options-container" v-if="optionsOpen">
    <button class="user_header_card-container open" @click="toggleOptions">
      <div class="user_header_card-picture">
        <img :src="pictureURL" alt="" />
        <div class="user_header_card-feeling">
          <div>
            <span v-if="feeling === 0">😐</span>
            <span v-if="feeling === 1">😀</span>
            <span v-if="feeling === 2">😍</span>
            <span v-if="feeling === 3">😡</span>
            <span v-if="feeling === 4">😭</span>
          </div>
        </div>
      </div>
      <div class="user_header_card-name">{{ username }}</div>
    </button>
    <button @click="openMyProfile" class="user_header_card-option">
      My Profile
    </button>
    <button @click="logout" class="user_header_card-option">Logout</button>
  </div>
</template>

<style lang="scss">
.user_header_card-options-container {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  position: absolute;
  top: 10px;
  right: 25px;
  width: fit-content;
  height: fit-content;
  z-index: 2;
  @media screen and (max-width: 1000px) {
    right: 25px;
    padding: 10px 20px;
    background-color: white;
    border-radius: 10px;
  }
  .user_header_card-option {
    background-color: transparent;
    border: none;
    padding: 0;
    margin: 0;
    margin-top: 5px;
    font-family: inherit;
    font-size: 20px;
    font-weight: 600;
    color: white;
    cursor: inherit;
    text-decoration: inherit;
    transition: all 0.3s ease-in-out;
    @media screen and (max-width: 1000px) {
      color: black;
    }
    &:hover {
      cursor: pointer;
      text-decoration: underline;
    }
  }
}
.user_header_card-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  font-family: 'Courier New', Courier, monospace;
  color: white;
  font-size: 20px;
  font-weight: bold;
  background-color: transparent;
  cursor: pointer;
  border: none;
  transition: all 0.2s ease;
  text-decoration: none;
  &:hover {
    transform: scale(1.02);
    @media screen and (min-width: 1000px) {
      transform: scale(1);
    }
  }
  &.open {
    margin-bottom: 15px;
    @media screen and (max-width: 1000px) {
      color: black;
    }
  }
  .user_header_card-picture {
    display: flex;
    flex-direction: row;
    align-items: center;
    position: relative;
    width: 50px;
    height: 50px;
    img {
      object-fit: cover;
      overflow: hidden;
      width: 100%;
      height: 100%;
      border-radius: 50%;
    }
    .user_header_card-feeling {
      position: absolute;
      bottom: -5px;
      right: -5px;
    }
  }
  .user_header_card-name {
    max-width: 150px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin-left: 20px;
    @media screen and (max-width: 650px) {
      display: none;
    }
  }
}
</style>
