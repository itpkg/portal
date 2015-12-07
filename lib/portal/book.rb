module Portal
  module Book
    ROOT = "#{Rails.root}/tmp/books"

    module_function
    def names
      Dir["#{ROOT}/*.html"].map { |f| f[(ROOT.size+1)..-6] }
    end
  end
end