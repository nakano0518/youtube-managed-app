export const ROUTES = {
    GET: {
        POPULARS: '/api/popular',
        VIDEO: '/api/video/:id',
        RELATED: '/api/related/:id',
        SEARCH: '/api/search'
    },
    POST: {
        TOGGLE_FAVORITE: '/api/favorite/:id/toggle'
    }
};

export default ROUTES;