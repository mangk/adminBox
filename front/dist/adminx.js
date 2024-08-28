window.adminX = {
    Name: 'adminX',
    Host: 'http://127.0.0.1:8910',
    IsRewriteIndex: false,
    /*** .writeByAdminX_config ***/
}

if (window.adminX.IsRewriteIndex && !(localStorage.getItem(_tokenStorageKey) && localStorage.getItem(_userTypeStroageKey) && localStorage.getItem(_userIdStroageKey))) {
    window.location.href = window.location.origin
}

/*** .writeByAdminX_func ***/