<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-4">
        <input v-model="new_kanban.title" placeholder="title" class="form-control">
        <input v-model="new_kanban.body" placeholder="body" class="form-control">
        <button @click="createKanban()" class="btn btn-success">カンバン作成</button>
      </div>
    </div>
  </div>
</template>
<script setup>
  class Kanban {
    constructor(title='', body='') {
      this.title = title
      this.body = body
      this.label = '未着手'
    }
  }

  const new_kanban = ref(new Kanban())

  const kanban_list = ref({
    waiting: [],
    working: [],
    completed: []
  })

  const createKanban = () => {
    const k = new_kanban.value
    if (k.title == '' || k.body =='') {
      return
    }
    kanban_list.value.waiting.push(k)
    new_kanban.value = new Kanban()
  }

</script>