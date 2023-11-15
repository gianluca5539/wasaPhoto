<script>
export default {
  name: 'HomeHeaderComponent',
  props: {
    name: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    pictureURL: {
      type: String,
      required: true
    }
  }
}
</script>

<template>
  <div class="homepage-header-container">
    <div class="homepage-header-title">WASAPhoto</div>
    <RouterLink to="/profile" class="homepage-header-profile-container">
      <div class="homepage-header-profile-picture">
        <img :src="pictureURL" alt="" />
        <div class="homepage-header-profile-feeling">
          <div>
            <span v-if="feeling === 0">😐</span>
            <span v-if="feeling === 1">😀</span>
            <span v-if="feeling === 2">😍</span>
            <span v-if="feeling === 3">😡</span>
            <span v-if="feeling === 4">😭</span>
          </div>
        </div>
      </div>
      <div class="homepage-header-profile-name">{{ name }}</div>
    </RouterLink>
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
  padding: 10px 30px;
  position: absolute;
  top: 0px;
  right: 0px;
  .homepage-header-title {
    font-size: 45px;
    font-weight: bold;
    color: white;
  }

  .homepage-header-profile-container {
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
    transition: all 0.5s ease;
    text-decoration: none;
    &:hover {
      transform: scale(1.05);
    }
    .homepage-header-profile-picture {
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-right: 20px;
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
      .homepage-header-profile-feeling {
        position: absolute;
        bottom: -5px;
        right: -5px;
      }
    }
    .homepage-header-profile-name {
      max-width: 150px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
}
</style>
