<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-4">
        <input v-model="new_kanban.title" placeholder="title" class="form-control">
        <input v-model="new_kanban.body" placeholder="body" class="form-control">
        <button @click="createKanban()" class="btn btn-success">カンバン作成</button>
      </div>
    </div>

    <hr>

    <div class="row mt-3">
      <div v-for="kanban_category in kanban_list" class="col-4 px-3">
        <div v-for="k in kanban_category" class="card">

          <div class="card-body">
            <h5 class="card-title">{{ k.title }}</h5>
            <p class="card-text">{{ k.body }}</p>
            <div class="d-flex justify-content-between">
              <button v-if="k.category >= 1" @click="moveKanban(k, 'left')" class="btn btn-dark">◀</button>
              <button v-else class="btn btn-secondary" disabled>◀</button>

              <button v-if="k.category <= 1" @click="moveKanban(k, 'right')" class="btn btn-dark">▶</button>
              <button v-else class="btn btn-secondary" disabled>▶</button>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
  class Kanban {
    constructor(title='', body='', category=0) {
      this.title = title
      this.body = body
      this.category = category
    }
  }

  const new_kanban = ref(new Kanban())

  const kanban_list = ref([
    [new Kanban('sample_waiting', 'sample_body')],
    [new Kanban('sample_working', 'sample_body', 1)],
    [new Kanban('sample_completed', 'sample_body', 2)]
  ])

  const createKanban = () => {
    const k = new_kanban.value
    if (k.title == '' || k.body =='') {
      return
    }
    kanban_list.value[0].push(k)
    new_kanban.value = new Kanban()
  }

  const moveKanban = (k, direction) => {
    // kanban_list内の対応するオブジェクトを取得
    const idx = kanban_list.value[k.category].indexOf(k)
    const kanban = kanban_list.value[k.category][idx]

    // 元居たカテゴリから削除
    kanban_list.value[kanban.category].splice(idx, 1)

    // カテゴリ移動
    if (direction == 'left') {
      kanban.category--
    } else {
      kanban.category++
    }
    kanban_list.value[kanban.category].push(kanban)
  }
</script>