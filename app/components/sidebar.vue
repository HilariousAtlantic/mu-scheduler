<template>

  <div class="sidebar">

    <slot></slot>

    <div class="sidebar-footer">

      <a href="https://github.com/hilariousatlantic">Developed by Hilarious Atlantic</a>

    </div>


  </div>

</template>

<script>

  export default {

    name: 'sidebar'

  }

</script>

<style lang="postcss" scoped>

  .sidebar {

    display: flex;
    flex-direction: column;
    background: #efefef;
    border-right: 1px solid #ddd;

  }

  .sidebar-content {

    padding: 20px;
    flex: 1;
    overflow-y: auto;

  }

  .sidebar-footer {

    text-align: center;
    padding: 10px 20px;
    border-top: 1px solid #ddd;

    a {

      color: #888;
      text-decoration: none;

      &:hover {

        color: #555;

      }

    }

  }

</style>
