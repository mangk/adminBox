window.adminX = {
    Name: 'adminX',
    Host: 'http://127.0.0.1:8910',
    IsRewriteIndex: false,
    /*** .writeByAdminX_config ***/
}

if (window.adminX.IsRewriteIndex && !(localStorage.getItem('x-token') && localStorage.getItem('x-user-type') && localStorage.getItem('x-user-id'))) {
    window.location.href = window.location.origin
}

/*** .writeByAdminX_func ***/