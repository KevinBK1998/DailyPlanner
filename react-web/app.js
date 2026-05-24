const form = document.querySelector('form');
const titleInput = document.querySelector('#task-title');
const taskList = document.querySelector('.task-list');
let nextTaskId = 4;

form.addEventListener('submit', (event) => {
    event.preventDefault();

    const title = titleInput.value.trim();
    if (!title) {
        return;
    }

    const taskId = `task-${nextTaskId}`
    const item = document.createElement('li');
    item.className = 'task-item pending';
    item.innerHTML = `
        <input id="${taskId}" type="checkbox" aria-label="Mark task as complete" />
        <div class="task-content">
            <label class="task-title" for="${taskId}">${title}</label>
            <span class="task-status">Pending</span>
        </div>
    `;
    taskList.appendChild(item);
    nextTaskId++;
    form.reset();
    titleInput.focus();
})