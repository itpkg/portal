require 'sidekiq/web'

Rails.application.routes.draw do

  resources :attachments, except: [:update, :edit, :new, :show]

  namespace :cms do
    resources :users, only:[:show, :index]
    resources :tags
    resources :articles do
      post :top
    end
    resources :comments, except: [:new, :show]
  end

  resources :notices, only: [:index, :create, :destroy]

  %w(info seo captcha adverts users about_us map).each do |a|
    get "site/#{a}"
    post "site/#{a}"
  end
  put 'site/role'
  get 'site/status'
  get 'site' => 'site#index'

  get 'ueditor' =>'ueditor#index'
  post 'ueditor' =>'ueditor#index'

  get 'personal/logs'

  resources :books, only: [:index, :show]

  get 'about_us' => 'home#about_us'
  get 'search' => 'home#search'
  #get ':sitemap.xml.gz' => 'home#sitemap', format: false, as: :sitemap
  get 'rss.atom' => 'home#rss', format: false, as: :rss
  get 'robots.txt' => 'home#robots', format: false, as: :robots
  get 'baidu_verify_(*id)', to: 'home#baidu', as: :baidu, format: true, constraints: {format: :html}
  get 'google(*id)', to: 'home#google', as: :google, format: true, constraints: {format: :html}


  namespace :questionnaire do
    resources :answers, only: [:destroy]
    resources :questions, except: [:index, :show]
    resources :reports do
      post :answer
      get :result
    end
  end

  resources :friend_links, only:[:index, :create, :destroy]
  resources :leave_words, only:[:index, :create, :destroy]
  get 'utils' =>'utils#index'

  devise_for :users

  authenticate :user, lambda { |u| u.is_admin? } do
    mount Sidekiq::Web => '/sidekiq'
  end



  # The priority is based upon order of creation: first created -> highest priority.
  # See how all your routes lay out with "rake routes".

  # You can have the root of your site routed with "root"
  root 'home#index'

  # Example of regular route:
  #   get 'products/:id' => 'catalog#view'

  # Example of named route that can be invoked with purchase_url(id: product.id)
  #   get 'products/:id/purchase' => 'catalog#purchase', as: :purchase

  # Example resource route (maps HTTP verbs to controller actions automatically):
  #   resources :products

  # Example resource route with options:
  #   resources :products do
  #     member do
  #       get 'short'
  #       post 'toggle'
  #     end
  #
  #     collection do
  #       get 'sold'
  #     end
  #   end

  # Example resource route with sub-resources:
  #   resources :products do
  #     resources :comments, :sales
  #     resource :seller
  #   end

  # Example resource route with more complex sub-resources:
  #   resources :products do
  #     resources :comments
  #     resources :sales do
  #       get 'recent', on: :collection
  #     end
  #   end

  # Example resource route with concerns:
  #   concern :toggleable do
  #     post 'toggle'
  #   end
  #   resources :posts, concerns: :toggleable
  #   resources :photos, concerns: :toggleable

  # Example resource route within a namespace:
  #   namespace :admin do
  #     # Directs /admin/products/* to Admin::ProductsController
  #     # (app/controllers/admin/products_controller.rb)
  #     resources :products
  #   end
end
