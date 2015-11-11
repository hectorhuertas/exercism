class Complement
  VERSION = 2
  def self.of_dna(word)
    convert_word(word, :dna)
  end

  def self.of_rna(word)
    convert_word(word, :rna)
  end

  def self.convert_word(word, type)
    word.chars.map do |letter|
      convert_letter(letter, type)
    end.join
  end

  def self.transformations
    { dna: { 'C' => 'G', 'G' => 'C', 'T' => 'A', 'A' => 'U' },
      rna: { 'G' => 'C', 'C' => 'G', 'A' => 'T', 'U' => 'A' } }
  end

  def self.convert_letter(letter, type)
    transformations[type][letter] || (fail ArgumentError)
  end
end
