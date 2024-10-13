window.adminBox = {
    Name: 'adminBox',
    Host: 'http://127.0.0.1:8910',
    IsRewriteIndex: false,
    BackendRouterPrefix: '',
    Locale: "zhCn",
    /*** .writeByadminBox_config ***/
}

if (window.adminBox.IsRewriteIndex && !(localStorage.getItem('x-token') && localStorage.getItem('x-user-type') && localStorage.getItem('x-user-id'))) {
    window.location.href = window.location.origin
}

/*** .writeByadminBox_func ***/