class BooksController < ApplicationController
  layout 'cms'

  ROOT = "#{Rails.root}/tmp/books"

  def show
    @title = params[:id]
    name = "#{ROOT}/#{@title}.html"
    if File.exists?(name)
      @body = File.read name
    else
      head :not_found
    end
  end

  def index
    @books = Dir["#{ROOT}/*.html"].map { |f| f[(ROOT.size+1)..-6] }
  end
end