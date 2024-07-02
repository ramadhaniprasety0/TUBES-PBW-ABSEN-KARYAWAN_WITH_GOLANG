document.addEventListener('DOMContentLoaded', function() {
    const adminList = document.getElementById('adminlist');

    function fetchAdmins() {
        fetch('/api/admin')
            .then(response => response.json())
            .then(data => {
                console.log('Data fetched from API:', data); // data yang diterima dari API /api/admin
                adminList.innerHTML = '';
                if (data.users && data.users.length > 0) {
                    data.users.forEach(user => {
                        const row = document.createElement('tr');
                        row.innerHTML = `
                            <td class="">${user.id}</td>
                            <td class="">${user.nama_admin}</td>
                            <td class="align-middle">
                                <button class="btn btn-primary btn-circle btn-sm" onclick="editAdmin(${user.id}, '${user.nama_admin}')"><i class="fa fa-pen-to-square"></i></button>
                                <button class="btn btn-danger btn-circle btn-sm" onclick="tampilHapusModal(${user.id}, '${user.nama_admin}')"><i class="fa-regular fa-trash-can"></i></button>
                            </td>
                        `;
                        adminList.appendChild(row);
                    });
                } else {
                    console.log('Admin Tidak ada');
                }
            })
            .catch(error => console.error('Error fetching admins:', error));
    }

    fetchAdmins();

    function updateClock() {
        const now = new Date();
        const date = now.toLocaleDateString('id-ID', {
            weekday: 'long',
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        });
        const time = now.toLocaleTimeString('id-ID', {
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit'
        });

        document.getElementById('date').value = date;
        document.getElementById('time').value = time;
    }

    setInterval(updateClock, 1000);
    updateClock();
});
